<template>
    <div class="file-manager-pane">
        <ul class="selected-file-list">
            <li v-for="f in files">
                <font-awesome-icon :icon="icons.file" size="lg"
                                   :class="{'html-file':f.type==='text/html','css-file':f.type==='text/css'}"/>
                <span>{{f.name}}</span>
                <font-awesome-icon :icon="icons.close" class="close-icon" size="1x" @click="removeFile(f)"/>
            </li>
        </ul>
    </div>
</template>

<script>
    import {faFile, faTimes} from '@fortawesome/free-solid-svg-icons';
    import {mapGetters, mapMutations} from 'vuex';

    export default {
        data() {
            return {
                icons: {
                    close: faTimes,
                    file: faFile,
                }
            }
        },
        computed: {
            ...mapGetters(['files'])
        },
        methods: {
            ...mapMutations(['removeFile'])
        }
    }
</script>

<style scoped>


</style>