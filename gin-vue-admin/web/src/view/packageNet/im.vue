<template>
    <div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            </div>
            <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
                <el-table-column type="selection" width="55" />
                <el-table-column align="left" label="ID" prop="ID" width="50" />
                <el-table-column align="left" label="平台名" prop="projectName" width="200" />
                <el-table-column align="left" label="服务器IP" prop="serverip" width="150" />
                <el-table-column align="left" label="JSON配置地址" prop="jsonconfig" width="370" />
                <el-table-column align="left" label="备注" prop="remark" width="370" />
                <el-table-column align="left" label="按钮组" min-width="160">
                    <template #default="scope">
                        <el-button type="primary" link icon="edit" @click="openImLineDialog(scope.row)">线路</el-button>
                        <el-button type="primary" link icon="edit" @click="updateim(scope.row)">变更</el-button>
                        <el-popover v-model="scope.row.visible" placement="top" width="160">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin-top: 8px;">
                                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" @click="deleteim(scope.row)">确定</el-button>
                            </div>
                            <template #reference>
                                <el-button type="primary" link icon="delete"
                                    @click="scope.row.visible = true">删除</el-button>
                            </template>
                        </el-popover>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>
        </div>
        <!-- form dialog -->
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="im">
            <el-form :model="form" label-width="80px">
                <el-form-item label="平台名称">
                    <el-input v-model="form.projectName" autocomplete="off" />
                </el-form-item>
                <el-form-item label="服务器ip">
                    <el-input v-model="form.serverip" autocomplete="off" />
                </el-form-item>
                <el-form-item label="JSON配置地址">
                    <el-input v-model="form.jsonconfig" autocomplete="off" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input type="textarea" v-model="form.remark" rows="3"></el-input>
                </el-form-item>

            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-dialog>

        <!-- show line dialog -->
        <el-dialog v-model="dialogTableVisible" :before-close="closeImLineDialog" title="im线路表" width="70%">
            <warning-bar title=" 线路的删除 与 添加 都是对浏览器本地数据进行修改 , save后才对json文件修改。" />
            <dev>
                <el-form :inline="true" :model="formInline" class="demo-form-inline">
                    <el-form-item label="线路域名">
                        <el-input v-model="formInline.base_url" placeholder="线路域名" />
                    </el-form-item>
                    <el-form-item label="资源域名">
                        <el-input v-model="formInline.res_url" placeholder="资源域名" />
                    </el-form-item>
                    <el-form-item label="socket ip">
                        <el-input v-model="formInline.socket_ip" placeholder="socket ip" />
                    </el-form-item>
                    <el-form-item label="port">
                        <el-select v-model="formInline.socket_port" placeholder="port">
                            <el-option label="9326" :value="parseInt('9326')" />
                            <el-option label="9325" :value="parseInt('9325')" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="time out">
                        <el-input v-model="formInline.timeout" placeholder="time out" />
                    </el-form-item>
                    <el-form-item label="ssl">
                        <el-select v-model="formInline.ssl" placeholder="ssl">
                            <el-option label="1" :value="parseInt('1')" />
                            <el-option label="2" :value="parseInt('2')" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="type">
                        <el-select v-model="formInline.type" placeholder="type">
                            <el-option label="0" :value="parseInt('0')" />
                            <el-option label="1" :value="parseInt('1')" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="备注">
                        <el-input v-model="formInline.remark" placeholder="备注" />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="addline()">添加</el-button>
                    </el-form-item>
                </el-form>
            </dev>
            <el-divider />
            <el-table :data="lines" :edit-config="{ editMethod: 'cell' }">
                <el-table-column property="base_url" label="线路" width="220">
                </el-table-column>
                <el-table-column property="res_url" label="资源线路" width="220" />
                <el-table-column property="socket_ip" label="socket ip" />
                <el-table-column property="socket_port" label="socket port" />
                <el-table-column property="timeout" label="timeout" />
                <el-table-column property="ssl" label="ssl" />
                <el-table-column property="remark" label="备注" />
                <el-table-column property="type" label="type" />
                <el-table-column fixed="right" label="Operations">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click="deleteLine(scope.$index)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-divider content-position="left">修改完记得保存</el-divider>
            <div>
                <el-button type="primary" @click="saveLines()">save</el-button>
            </div>
            <!-- <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template> -->
        </el-dialog>

    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { getImList, createIm, updateIm, deleteIm, getImById } from "../../api/im"
import { getImLine, changeLine } from '../../api/imlinesapi'
import { ElMessage } from 'element-plus'
import { Line } from '../../interface/imline'
import WarningBar from '../../components/warningBar/warningBar.vue'

const form = ref({
    projectName: "",
    serverip: "",
    jsonconfig: "",
    remark: "",
})



const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const getTableData = async () => {
    const data = await getImList()
    tableData.value = data.data
}
getTableData()


// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    // getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    // getTableData()
}

const closeDialog = () => {
    dialogFormVisible.value = false
    form.value = {
        projectName: "",
        serverip: "",
        jsonconfig: "",
        remark: "",
    }
}

const dialogFormVisible = ref(false)
const type = ref('')

const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

const enterDialog = async () => {
    let res
    switch (type.value) {
        case 'create':
            console.log(form.value)
            res = await createIm(form.value)
            break
        case 'update':
            res = await updateIm(form.value)
            break
        default:
            res = await createIm(form.value)
            break
    }

    if (res.code === 0) {
        closeDialog()
        getTableData()
    }
}

const deleteim = async (row) => {
    row.visible = false
    const res: any = await deleteIm({ ID: row.ID })

    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '删除成功'
        })
        if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

const updateim = async (row) => {
    const res: any = await getImById({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        console.log(res.data)
        form.value = res.data
        dialogFormVisible.value = true
    }
}

const dialogTableVisible = ref(false)
const formInline = ref<Line>({
    base_url: "",
    res_url: "",
    socket_ip: "",
    socket_port: 9326,
    timeout: 30000,
    ssl: 2,
    remark: "",
    type: 0
})
const lines = ref<Line[]>([])
const imid = ref<string>()
const openImLineDialog = async (row) => {
    // type.value = 'create'
    dialogTableVisible.value = true

    const data = await getImLine({ ID: row.ID })
    lines.value = data.data
    imid.value = <string>row.ID
    formInline.value = {
        timeout: 30000,
        res_url: lines.value[0].res_url,
        socket_ip: lines.value[0].socket_ip
    }
}

const closeImLineDialog = async () => {
    dialogTableVisible.value = false
}

const addline = async () => {
    lines.value?.push(formInline.value)
    console.log(lines.value)
    formInline.value = {
        timeout: 30000,
        ssl: 2,
        type: 0,
    }
}

const deleteLine = async (index: number) => {
    console.log(index)
    lines.value?.splice(index, 1)
}


const saveLines = async () => {
    console.log(lines.value)
    console.log(imid.value)
    const data: any = await changeLine(lines.value, <string>imid.value)
    if (data.code === 0) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
    }
}
</script>

<script  lang="ts">
export default {
    name: 'im'
}
</script>

<style scoped></style>