<template>
    <div>
        <MenuBar></MenuBar>
        <FileManager></FileManager>
        <AnalyzedData></AnalyzedData>
    </div>
</template>

<script>
    import MenuBar from '@/components/MenuBar.vue';
    import FileManager from '@/components/FileManager.vue';
    import AnalyzedData from '@/components/AnalyzedData.vue';

    export default {
        components: {
            MenuBar,
            FileManager,
            AnalyzedData,
        }
    }
</script>

<style lang="scss">

</style>
