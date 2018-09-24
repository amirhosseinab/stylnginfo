<template>
    <ul class="file-list">
        <li v-for="f in files" class="file-list-item">
            <div class="file">
                <font-awesome-icon :icon="icons.file" size="lg"
                                   :class="{'html':f.type==='text/html','css':f.type==='text/css'}"/>
                <span class="file-name">{{f.name}}</span>
                <font-awesome-icon :icon="icons.close" class="remove icon" size="1x" @click="removeFile(f)"
                                   :class="{'disabled':waiting}" v-show="!analyzed"/>
            </div>
        </li>
    </ul>
</template>
<script>
    import {faFile, faTimes} from '@fortawesome/free-solid-svg-icons';
    import {mapGetters,mapMutations} from 'vuex';

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
            ...mapGetters(['files', 'analyzed', 'waiting'])
        },
        methods: {
            ...mapMutations(['removeFile'])
        }
    }
</script>
<style lang="scss">
    ul.file-list {
        display: block;
        margin: 0;
        padding: 0;
        text-align: left;
        .file-list-item {
            text-align: left;
            padding: 6px 10px;
            display: block;
            font-family: monospace;
            font-size: 1.2em;
            cursor: default;
            &:hover {
                background-color: lighten($dark-gray-color, 7%);
            }
            .file {
                .file-name {
                    color: $light-gray-color;
                    margin-left: 7px;
                    display: inline-block;
                }
                .html {
                    color: $blue-color;
                }
                .css {
                    color: $green-color;
                }
                .remove.icon {
                    color: lighten($red-color, 15%);
                    cursor: pointer;
                    float: right;
                    display: inline-block;
                    height: 20px;
                    width: 20px;
                    padding: 4px
                }
            }
        }
    }

</style>