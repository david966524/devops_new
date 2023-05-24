<template>
    <div>
        <div class="gva-table-box">
            <div class="card-container">
                <el-card style="{width: 100% !important}" shadow="never">
                    <div>
                        <el-row>
                            <el-form :model="form">
                                <el-form-item label="域名">
                                    <el-input v-model="form.url" />
                                </el-form-item>
                            </el-form>

                            <el-button :icon="Search" @click="checkDomain" circle />
                        </el-row>
                    </div>
                    <div slot="header" class="clear-fix">
                        <i class="float-right el-icon-arrow-down"></i>
                        <el-descriptions title="Info" column="1"  >
                            <el-descriptions-item label="域名"  >{{ result.url }}</el-descriptions-item>
                            <el-descriptions-item label="QQ"> {{ result.qq_msg }}</el-descriptions-item>
                            <el-descriptions-item label="微信">{{ result.vx_msg }}</el-descriptions-item>
                            <el-descriptions-item label="原因">{{ result.cause }}</el-descriptions-item>
                            <el-descriptions-item label="icp">{{ result.icp }}</el-descriptions-item>
                            <el-descriptions-item label="icp 信息">{{ result.icp_name }}</el-descriptions-item>
                        </el-descriptions>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { CheckDomain } from "../../../api/chenckdomain"
import { ElMessage } from 'element-plus'
import {
    Check,
    Delete,
    Edit,
    Message,
    Search,
    Star,
} from '@element-plus/icons-vue'

const form = ref({
    url: ""
})

const result = ref({
    url: "",
    qq_msg: "",
    vx_msg: "",
    cause: "",
    icp_name: "",
    icp: ""
})

const checkDomain = async () => {
    const data = await CheckDomain({
        Url: form.value.url,
    })
    result.value = data.data
}
</script>

<script  lang="ts">
export default {
    name: 'checkDomain'
}
</script>

<style scoped>

</style>