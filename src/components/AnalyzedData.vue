<template>
    <div class="result-pane">
        <div v-show="waiting" class="waiting-block">
            <font-awesome-icon :icon="icons.waiting" size="2x" spin/>
            <span>Please wait...</span>
        </div>

        <div v-for="hf in htmlFiles" class="html-file">
            <div>{{hf.name}}</div>
            <div class="css-files-block">
                <div v-for="cf in hf.cssFiles" class="css-file">
                    {{cf.name}}
                    <div class="selectors-block">
                        <div v-for="s in cf.selectors" class="selector"
                             @click="selectSelectors({'htmlFileName':hf.name,'selector':s.name})"
                             :class="{'selected':s.selected}">
                            {{s.name}}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapMutations} from 'vuex'
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
            ...mapGetters(['htmlFiles', 'waiting'])
        },
        methods: {
            ...mapMutations(['selectSelectors'])
        }
    }
</script>

<style scoped>

</style>