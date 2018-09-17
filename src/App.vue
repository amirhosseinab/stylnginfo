<template>
    <div>
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
