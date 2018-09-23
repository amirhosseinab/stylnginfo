<template>
    <div class="result-pane">
        <div v-show="waiting" class="waiting-block">
            <font-awesome-icon :icon="icons.waiting" size="2x" spin/>
            <span>Please wait...</span>
        </div>
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
    import {faSpinner} from '@fortawesome/free-solid-svg-icons'

    export default {
        data() {
            return {
                icons: {
                    waiting: faSpinner,
                }
            }
        },
        computed: {
            ...mapGetters(['analyzedData', 'relatedCssFiles', 'relatedSelectors', 'waiting'])
        }
    }
</script>

<style scoped>

</style>