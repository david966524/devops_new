<template>
    <div>
        <warning-bar title="啦啦啦啦啦~" />
        <div class="gva-table-box">
            <!-- <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            </div> -->
            <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
                <el-table-column type="selection" width="55" />
                <el-table-column align="left" label="ID" prop="domainsID" width="180" />
                <el-table-column align="left" label="域名" prop="domains" width="200" />
                <el-table-column align="left" label="状态" prop="state" width="120">
                    <template #default="scope">
                        <el-tag :type="scope.row.state === 0 ? 'success' : 'danger'" disable-transitions>{{
                            scope.row.state ===
                            0 ? '正常' : '不正常'
                        }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="用户锁定" prop="userLock" width="120" />
                <el-table-column align="left" label="管理员锁定" prop="adminLock" width="120" />
                <el-table-column align="left" label="按钮组" min-width="160">
                    <template #default="scope">
                        <el-button type="primary" link icon="View" @click="showrecord(scope.row)">查看解析</el-button>
                        <el-button type="primary" link icon="edit" @click="addrecord(scope.row)">添加解析</el-button>
                        <el-button type="primary" link icon="edit" @click="autoaddrecord(scope.row)">一键添加解析</el-button>
                        <!-- <el-popover v-model="scope.row.visible" placement="top" width="160">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin-top: 8px;">
                                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" @click="deleteCustomer(scope.row)">确定</el-button>
                            </div>
                            <template #reference>
                                <el-button type="primary" link icon="delete"
                                    @click="scope.row.visible = true">删除</el-button>
                            </template>
                        </el-popover> -->
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>
        </div>
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="解析表单">
            <el-form :inline="true" :model="form" label-width="80px">
                <el-form-item label="记录类型">
                    <el-select v-model="form.type" class="m-2" placeholder="Select">
                        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="主机头">
                    <el-input v-model="form.host" autocomplete="off" />
                </el-form-item>
                <el-form-item label="记录值">
                    <el-input v-model="form.value" autocomplete="off" />
                </el-form-item>
                <el-form-item label="TTL">
                    <el-input v-model="form.TTL" autocomplete="off" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="form.remark" autocomplete="off" />
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-dialog>


        <el-dialog v-model="dialogtableRecord" :before-close="closetableRecordDialog" title="解析表" width="70%">
            <el-table :data="recordtable" :edit-config="{ editMethod: 'cell' }">
                <el-table-column property="record" label="record" width="120" />
                <el-table-column property="type" label="type" width="120" />
                <el-table-column property="value" label="value" width="300" />
                <el-table-column property="TTL" label="TTL" width="120" />
                <el-table-column property="state" label="state" width="120" />
                <el-table-column property="remark" label="remark" width="200" />
                <!-- <el-table-column fixed="right" label="Operations">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click="">删除</el-button>
                    </template>
                </el-table-column> -->
            </el-table>
        </el-dialog>

        <el-dialog v-model="dialogautoRecord" :before-close="closeautoRecordDialog" title="一键添加解析" width="70%">
            <warning-bar title="除了 TTl 备注 都是必填项" />
            <el-form :inline="true" :model="autoform" label-width="80px">
                <el-form-item label="记录类型">
                    <el-select v-model="autoform.type" class="m-2" placeholder="Select">
                        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="主机头">
                    <el-input v-model="autoform.host" autocomplete="off" />
                </el-form-item>
                <el-form-item label="记录值">
                    <el-select v-model="autoform.value" class="m-2" placeholder="Select">
                        <el-option v-for="item in vhost" :key="item.value" :label="item.label" :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="im">
                    <el-select v-model="autoform.imIp" class="m-2" placeholder="Select">
                        <el-option v-for="item in im" :key="item.projectName" :label="item.projectName"
                            :value="item.serverip" />
                    </el-select>
                </el-form-item>
                <el-form-item label="TTL">
                    <el-input v-model="autoform.TTL" autocomplete="off" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="autoform.remark" autocomplete="off" />
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeautoRecordDialog">取 消</el-button>
                    <el-button type="primary" @click="autoAddim">确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>
  
<script setup lang="ts">
import {
    createExaCustomer,
    updateExaCustomer,
    deleteExaCustomer,
    getExaCustomer,
    getExaCustomerList
} from '../../api/customer'
import { getdnsList, getdnsrecord, adddnsrecord, addautodnsrecord } from '../../api/dnsapi'
import WarningBar from '../../components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '../../utils/format'
import { getImList } from '../../api/im'
//添加解析 form
const form = ref({
    domainID: "",
    type: "",
    host: "",
    value: "",
    TTL: "300",
    remark: "",
})
//一键添加解析 form
const autoform = ref({
    domainID: "",
    domainName: "",
    type: "CNAME",
    host: "",
    value: "",
    TTL: "300",
    remark: "",
    imIp: ""
})
// 添加解析 form select
const options = [
    {
        value: 'CNAME',
        label: 'CNAME',
    },
    {
        value: 'A',
        label: 'A',
    },
    {
        value: 'TXT',
        label: 'TXT',
    },
]
// 解析表对象
const recordtable = ref([])
// 解析表 dialog 开关
const dialogtableRecord = ref(false)

// 分页对象
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
    const table: any = await getdnsList({ page: page.value, pageSize: pageSize.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()
//
const dialogFormVisible = ref(false)
const type = ref('')
const updateCustomer = async (row) => {
    const res = await getExaCustomer({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        form.value = res.data.customer
        dialogFormVisible.value = true
    }
}
// 关闭 dialog
const closeDialog = () => {
    dialogFormVisible.value = false
    form.value = {
        domainID: "",
        type: "",
        host: "",
        value: "",
        TTL: "300",
        remark: "",
    }
}
// 关闭 解析表 dialog
const closetableRecordDialog = () => {
    dialogtableRecord.value = false
}
const deleteCustomer = async (row) => {
    row.visible = false
    const res = await deleteExaCustomer({ ID: row.ID })
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
// dialog 入口
const enterDialog = async () => {
    let res
    switch (type.value) {
        case 'create':
            console.log(form.value)
            res = await adddnsrecord(form.value)
            break
        case 'update':
            res = await updateExaCustomer(form.value)
            break
        default:
            res = await adddnsrecord(form.value)
            break
    }

    if (res.code === 0) {
        closeDialog()
        getTableData()
    }
}
// 添加解析
const addrecord = async (row) => {
    type.value = 'create'
    dialogFormVisible.value = true
    //await adddnsrecord({})
    form.value.domainID = row.domainsID
}
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}
// 解析表
const showrecord = async (row) => {
    dialogtableRecord.value = true
    const data = await getdnsrecord({
        DomainID: row.domainsID
    })
    recordtable.value = data.data
}
// 一键添加解析
const dialogautoRecord = ref(false)
const vhost = ref([
    {
        value: 'my101-site-03.cdn-ng.net.',
        label: 'my101-site-03',
    },
    {
        value: 'my101-site-05.cdn-ng.net.',
        label: 'my101-site-05',
    },
])
const im: any = ref([])
// 打开自动添加域名dailog
const autoaddrecord = async (row) => {
    dialogautoRecord.value = true
    autoform.value.domainID = row.domainsID
    autoform.value.domainName = row.domains
    const data = await getImList()
    im.value = data.data
}
// 关闭自动添加域名dailog 重置表单
const closeautoRecordDialog = () => {
    dialogautoRecord.value = false
    autoform.value = {
        domainID: "",
        domainName: "",
        type: "CNAME",
        host: "",
        value: "",
        TTL: "300",
        remark: "",
        imIp: "",
    }
}
// 提交自动添加域名表单
const autoAddim = async () => {
    console.log(autoform.value)
    const data: any = await addautodnsrecord(autoform.value)
    if (data.code === 0) {
        closeautoRecordDialog()
        ElMessage({
            message: data.msg,
            type: 'success',
        })
    }
}

</script>
  
<script lang="ts">

export default {
    name: 'dns'
}
</script>
  
<style></style>
  