<template>
    <div>
        <MenuBar></MenuBar>
        <FileManager v-show="files.length"></FileManager>
        <AnalyzedData v-show="files.length"></AnalyzedData>
    </div>
</template>

<script>
    import MenuBar from '@/components/MenuBar.vue';
    import FileManager from '@/components/FileManager.vue';
    import AnalyzedData from '@/components/AnalyzedData.vue';
    import {mapGetters} from 'vuex'

    export default {
        computed: {
            ...mapGetters(['files'])
        },
        components: {
            MenuBar,
            FileManager,
            AnalyzedData,
        }
    }
</script>

<style lang="scss">

</style>
