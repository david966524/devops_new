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
                <el-table-column align="left" label="备注" prop="remark" width="370" />
                <el-table-column align="left" label="按钮组" min-width="250">
                    <template #default="scope">
                        <el-button type="primary" link icon="edit" @click="updateairborne(scope.row)">变更</el-button>
                        <el-popover v-model="scope.row.visible" placement="top" width="160">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin-top: 8px;">
                                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" @click="deleteairborne(scope.row)">确定</el-button>
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
      
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { getAirborneList, deleteAirborne, getAirborneById, updateAirborne, createAirborne } from "../../api/airborne"
import { ElMessage } from 'element-plus'

const form = ref({
    projectName: "",
    serverip: "",
    remark: "",
})


const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const getTableData = async () => {
    const data = await getAirborneList()
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
            res = await createAirborne(form.value)
            break
        case 'update':
            res = await updateAirborne(form.value)
            break
        default:
            res = await createAirborne(form.value)
            break
    }

    if (res.code === 0) {
        closeDialog()
        getTableData()
    }
}

const deleteairborne = async (row) => {
    row.visible = false
    const res: any = await deleteAirborne({ ID: row.ID })

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

const updateairborne = async (row) => {
    const res: any = await getAirborneById({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        console.log(res.data)
        form.value = res.data
        dialogFormVisible.value = true
    }
}



</script>

<script  lang="ts">
export default {
    name: 'im'
}
</script>

<style scoped></style>