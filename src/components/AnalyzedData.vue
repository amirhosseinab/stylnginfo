<template>
    <div class="result-pane">
        <template v-for="hf in analyzedData.htmlFiles">
            <div class="html-file">
                <div>{{hf.name}}</div>
                <div class="css-files-block">
                    <div v-for="cf in relatedCssFiles(hf.name)" class="css-file">
                        {{cf}}

                        <div class="selectors-block">
                            <div v-for="sel in relatedSelectors(hf.name,cf)" class="selector">
                                {{sel}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </template>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'

    export default {
        computed: {
            ...mapGetters(['analyzedData', 'relatedCssFiles','relatedSelectors'])
        }
    }
</script>

<style scoped>

</style>