<template>
    <div>
        <HeaderSection/>
        <Sidebar></Sidebar>
        <ResultPane></ResultPane>
    </div>
</template>

<script>
    import HeaderSection from '@/components/HeaderSection.vue';
    import Sidebar from '@/components/Sidebar.vue';
    import ResultPane from '@/components/ResultPane.vue';
    import {mapGetters} from 'vuex'

    export default {
        computed: {
            ...mapGetters(['files'])
        },
        components: {
            HeaderSection,
            Sidebar,
            ResultPane,
        }
    }
</script>
