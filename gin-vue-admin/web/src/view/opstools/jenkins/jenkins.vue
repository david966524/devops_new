<template>
    <div>
        <p>111</p>
    </div>
</template>

<script setup lang="ts">
import { getjobs } from "../../../api/jenkinsapi"

const getJobs = async () => {
    const data = await getjobs()
    console.log(data.data)
}
getJobs()
</script>

<style scoped></style>