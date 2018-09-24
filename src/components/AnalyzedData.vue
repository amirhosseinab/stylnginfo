<template>
    <div class="result-pane">
        <Waiting/>
        <!--<div v-for="hf in htmlFileNames" class="html-file">-->
            <!--<div>{{hf.name}}</div>-->
            <!--<div class="css-files-block">-->
                <!--<div v-for="cf in hf.cssFiles" class="css-file">-->
                    <!--{{cf.name}}-->
                    <!--<div class="selectors-block">-->
                        <!--<div v-for="s in cf.selectors" class="selector"-->
                             <!--@click="selectSelectors({'htmlFileName':hf.name,'selector':s.name})"-->
                             <!--:class="{'selected':s.selected}">-->
                            <!--{{s.name}}-->
                        <!--</div>-->
                    <!--</div>-->
                <!--</div>-->
            <!--</div>-->
        <!--</div>-->
    </div>
</template>

<script>
    import {mapGetters, mapMutations} from 'vuex'
    import Waiting from "@/components/Waiting";

    export default {
        components: {
            Waiting
        },
        computed: {
            ...mapGetters(['htmlFileNames'])
        },
        methods: {
            ...mapMutations(['selectSelectors'])
        }
    }
</script>

