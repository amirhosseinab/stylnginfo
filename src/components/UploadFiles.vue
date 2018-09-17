<template>
    <div>
        <input type="file" ref="fileInput" @change="handleFiles" multiple accept="text/css, text/html"/>
        <input type="button" value="Add HTML & CSS Files" @click="selectFile()">
        <ul class="selected-file-list">
            <li v-for="f in files">
                <font-awesome-icon :icon="icons.code" class="file-icon"/>
                <span>{{f.name}}</span>
                <font-awesome-icon :icon="icons.close" class="close-icon" size="s" @click="remove(f)"/>
            </li>
        </ul>
        <input type="button" value="Analyze" @click="uploadFiles" v-show="files.length">
    </div>
</template>

<script>
    import {faCode, faTimes} from '@fortawesome/free-solid-svg-icons'

    export default {
        data() {
            return {
                icons: {
                    close: faTimes,
                    code: faCode,
                },
                files: [],
                urls: []
            }
        },
        methods: {
            selectFile() {
                this.$refs.fileInput.click()
            },
            remove(f) {
                this.files.splice(this.files.indexOf(f), 1)
            },
            handleFiles(e) {
                let fileList = [];
                if (e.target.files) {
                    fileList.push(...e.target.files)
                }
                fileList.forEach(fl => {
                    if (!this.files.find(f => f.name === fl.name)) {
                        this.files.push(fl);
                    }
                });
            },
            uploadFiles() {

            }
        }
    }
</script>

<style scoped>
    .file-icon {
        color: #55a149;
        margin-right: 5px;
    }

    .close-icon {
        color: #ff6262;
        float: right;
        cursor: pointer;
    }
</style>