<template>
    <div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            </div>
            <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
                <el-table-column align="left" label="ID" prop="ID" width="120" />
                <el-table-column align="left" label="账户名" prop="Name" width="180" />
                <el-table-column align="left" label="Email" prop="Email" width="200" />
                <el-table-column align="left" label="api密钥" prop="Key" width="350" />
                <!-- <el-table-column align="left" label="状态" prop="Type" width="120" /> -->
                <el-table-column align="left" label="状态" prop="Type" width="120">
                    <template #default="scope">
                        <el-tag :type="scope.row.Type === 1 ? 'success' : 'danger'" disable-transitions>{{ scope.row.Type ===
                            1 ? '启用' : '禁用'
                        }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="按钮组" width="160">
                    <template #default="scope">
                        <el-button type="primary" link icon="edit" @click="updateCfApi(scope.row)">变更</el-button>
                        <el-popover v-model="scope.row.visible" placement="top" width="160">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin-top: 8px;">
                                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" @click="deletecfapitest(scope.row)">确定</el-button>
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
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="账户">
            <el-form :inline="true" :model="form" label-width="80px">
                <el-form-item label="账户名">
                    <el-input v-model="form.Name" autocomplete="off" />
                </el-form-item>
                <el-form-item label="email">
                    <el-input v-model="form.Email" autocomplete="off" />
                </el-form-item>
                <el-form-item label="apikey">
                    <el-input v-model="form.Key" autocomplete="off" />
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="form.Type">
                        <el-option label="启用" :value=parseInt(1)></el-option>
                        <el-option label="禁用" :value=parseInt(0)></el-option>
                    </el-select>
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
  
<script setup  >


import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { createCfApi, getCfApiList, deleteCfApi, getCfApiById ,updateCfApiForm} from '@/api/apimanager'



const form = ref({
    Name: "",
    Email: "",
    Key: "",
    Type: "",
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}



// 查询
const getTableData = async () => {
    const table = await getCfApiList()
    if (table.code === 0) {
        tableData.value = table.data
    }
}
getTableData()

const dialogFormVisible = ref(false)
const type = ref('')

const updateCfApi = async (row) => {
    const res = await getCfApiById({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        console.log(res.data)
        form.value = res.data
        dialogFormVisible.value = true
    }
}

const closeDialog = () => {
    dialogFormVisible.value = false
    form.value = {
        Name: "",
        Email: "",
        Key: "",
        Type: "",
    }
}


const deletecfapitest = async (row) => {
    row.visible = false
    const res = await deleteCfApi({ ID: row.ID })
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

const enterDialog = async () => {
    let res
    switch (type.value) {
        case 'create':
            console.log(form.value)
            res = await createCfApi(form.value)
            break
        case 'update':
            res = await updateCfApiForm(form.value)
            break
        default:
            res = await createCfApi(form.value)
            break
    }

    if (res.code === 0) {
        closeDialog()
        getTableData()
    }
}
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

</script>
  
<script >

export default {
    name: 'apikeymanager'
}
</script>
  
<style></style>
  