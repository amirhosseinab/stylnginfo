<template>
    <div>
        <HeaderSection/>
        <Sidebar></Sidebar>
        <AnalyzedData v-show="files.length"></AnalyzedData>
    </div>
</template>

<script>
    import HeaderSection from '@/components/HeaderSection.vue';
    import Sidebar from '@/components/Sidebar.vue';
    import AnalyzedData from '@/components/AnalyzedData.vue';
    import {mapGetters} from 'vuex'

    export default {
        computed: {
            ...mapGetters(['files'])
        },
        components: {
            HeaderSection,
            Sidebar,
            AnalyzedData,
        }
    }
</script>

<style lang="scss">

</style>
