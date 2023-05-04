<template>
    <div>
        <div class="gva-table-box">

            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            </div>
            <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
                <el-tab-pane v-for="(tab, index) in tabs" :key="index" :label="tab.name" :name="tab.name">{{ tab.name
                }}</el-tab-pane>
                <el-table :data="asiacloudtable" style="width: 100%">
                    <el-table-column prop="domain" label="domain" width="200" />
                    <el-table-column prop="host" label="host" width="200" />
                    <el-table-column prop="vhost" label="vhost" width="200" />
                    <el-table-column align="left" prop="status" label="SSLstatus" width="120">
                        <template #default="scope">
                            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'" disable-transitions>{{
                                scope.row.status ===
                                1 ? '开启' : '未启'
                            }}</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="cname" label="cname" width="400" />
                    <el-table-column fixed="right" label="Operations" width="200">
                        <template #default="scope">
                            <el-button type="primary" link icon="edit" @click="updateDomain(scope.row)">变更</el-button>
                            <el-popover v-model="scope.row.visible" placement="top" width="160">
                                <p>确定要删除吗？</p>
                                <div style="text-align: right; margin-top: 8px;">
                                    <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                                    <el-button type="primary" @click="deletedomain(scope.row)">确定</el-button>
                                </div>
                                <template #reference>
                                    <el-button type="primary" link icon="delete"
                                        @click="scope.row.visible = true">删除</el-button>
                                </template>
                            </el-popover>
                        </template>
                    </el-table-column>
                </el-table>
            </el-tabs>
        </div>
        <!-- form dialog -->
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="域名表单">
            <el-form :model="form" label-width="80px">
                <el-form-item  label="domain">
                    <el-input v-model="form.domain" autocomplete="off" />
                </el-form-item>
                <el-form-item label="host">
                    <el-input v-model="form.host" autocomplete="off" />
                </el-form-item>
                <el-form-item label="Zones" label-width="100">
                    <el-select v-model="form.vhost" placeholder="Please select a zone">
                        <el-option label="my101-site-03" value="my101-site-03" />
                        <el-option label="my101-site-04" value="my101-site-04" />
                        <el-option label="my101-site-05" value="my101-site-05" />
                    </el-select>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog(form.vhost)">确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { ElMessage } from 'element-plus'
import WarningBar from '../../components/warningBar/warningBar.vue'
import type { TabsPaneContext } from 'element-plus'
import { getAsiaCloudVhost, getAsiaCloudDomainList, deleteAsiaCloudDomain ,addAsiaCloudDomain ,updateAsiaCloudDomain} from '../../api/asiacloudapi'

const form = ref({
    id:0,
    domain: "",
    host: "",
    vhost:""
})
const asiacloudtable = ref()
const activeName = ref('first')
const tabs = ref()
const handleClick = async (tab: TabsPaneContext, event: Event) => {
    console.log(tab.paneName, event)
    // const data = await getAsiaCloudDomainList(<string>tab.paneName)
    // asiacloudtable.value = data.data
    getTableData(<string>tab.paneName)
}
const getVhost = async () => {
    const data = await getAsiaCloudVhost()
    tabs.value = data.data
}
 
getVhost()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const getTableData = async (paneName)=>{
    const data = await getAsiaCloudDomainList(<string>paneName)
    asiacloudtable.value = data.data
}

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
        id: 0,
        domain: "",
        host: "",
        vhost:""
    }
}

const dialogFormVisible = ref(false)
const type = ref('')

const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}



const deletedomain = async (row) => {
    row.visible = false
    console.log(row)
    const res: any = await deleteAsiaCloudDomain(row)
    if (res.code === 0) {
        getTableData(row.vhost)
        ElMessage({
            type: 'success',
            message: '删除成功'
        })
        if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
    }
}

const enterDialog = async (vhost) => {
    let res
    switch (type.value) {
        case 'create':
            console.log(form.value)
            res = await addAsiaCloudDomain(form.value)
            break
        case 'update':
            res = await updateAsiaCloudDomain(form.value)
            break
        default:
            res = await addAsiaCloudDomain(form.value)
            break
    }

    if (res.code === 0) {
        closeDialog()
        getTableData(vhost)
    }
}

const updateDomain = async (row) => {
    dialogFormVisible.value = true
    type.value = 'update'
    form.value = {
        id: row.id,
        domain: row.domain,
        host: row.host,
        vhost: row.vhost
    }
}


</script>

<script  lang="ts">
export default {
    name: 'asiacloud'
}
</script>

<style scoped></style>