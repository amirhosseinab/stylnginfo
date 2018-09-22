<template>
    <div class="result-pane">
        <ul>
            <li v-for="hf in analyzedData.htmlFiles" class="html-file">
                {{hf.name}}
                <div v-for="selector in hf.appliedSelectors">
                    <span class="css-file-name">{{selector.cssFileName}}</span>
                    <span class="css-selector-name">{{selector.selectorName}}</span>
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'

    export default {
        computed: {
            ...mapGetters(['analyzedData'])
        }
    }
</script>

<style scoped>

</style>