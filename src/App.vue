<template>
    <div>
        <h1>Please Select HTML & CSS Files to Analyze Them</h1>
        <UploadFiles></UploadFiles>
        <router-view/>
    </div>
</template>

<script>
    import UploadFiles from '@/components/UploadFiles.vue';

    export default {
        components: {
            UploadFiles
        }
    }
</script>

<style lang="scss">

</style>
