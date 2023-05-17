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
                        <el-descriptions title="User Info" column="1">
                            <el-descriptions-item label="Username">kooriookami</el-descriptions-item>
                            <el-descriptions-item label="Telephone">18100000000</el-descriptions-item>
                            <el-descriptions-item label="Place">Suzhou</el-descriptions-item>
                            <el-descriptions-item label="Remarks">
                                <el-tag size="small">School</el-tag>
                            </el-descriptions-item>
                            <el-descriptions-item label="Address">No.1188, Wuzhong Avenue, Wuzhong District, Suzhou,
                                Jiangsu
                                Province</el-descriptions-item>
                        </el-descriptions>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
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

const checkDomain = async () => {
    console.log(form.value.url)
}
</script>

<script  lang="ts">
export default {
    name: 'checkDomain'
}
</script>

<style scoped></style>