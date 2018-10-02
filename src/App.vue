<template>
    <div>
        <HeaderSection/>
        <Sidebar></Sidebar>
        <ResultPane></ResultPane>
        <footer>
            <span>
                made with <span class="heart">&hearts;</span> by <a href="https://github.com/amirhosseinab"
                                                                    title="Amir Hossein Abdolkhalegh"
                                                                    target="_blank">@amirhosseianb</a>
            </span>
        </footer>
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
