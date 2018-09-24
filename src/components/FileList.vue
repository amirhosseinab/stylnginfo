<template>
    <div>
        <div v-show="!files.length" class="info">
            HTML & CSS files will show here
        </div>
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
        max-height: $file-list-height;
        overflow: auto;
        border-bottom: solid 3px $border-color;
        @extend %scrollbar;
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
                    padding-left: 7px;
                    display: inline-block;
                    width: 87%;
                    vertical-align: middle;
                }
                .html, .css {
                    width: 5%;
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
                    display: inline-block;
                    height: 20px;
                    width: 8%;
                    min-width: 13px;
                    padding: 4px;
                    vertical-align: middle;
                }
            }
        }
    }

</style>