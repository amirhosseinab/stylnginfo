<template>
    <div>
        <input type="file" ref="fileInput" @change="handleFiles" multiple accept="text/css, text/html"/>
        <input type="button" value="Add HTML & CSS Files" @click="selectFile()">
        <ul class="selected-file-list">
            <li v-for="f in files">
                <font-awesome-icon :icon="icons.file" size="lg"
                                   :class="{'html-file':f.type==='text/html','css-file':f.type==='text/css'}"/>
                <span>{{f.name}}</span>
                <font-awesome-icon :icon="icons.close" class="close-icon" size="1x" @click="remove(f)"/>
            </li>
        </ul>
        <input type="button" value="Analyze" @click="uploadFiles" v-show="files.length">
    </div>
</template>

<script>
    import {faFile, faTimes} from '@fortawesome/free-solid-svg-icons';
    import axios from 'axios';

    export default {
        data() {
            return {
                icons: {
                    close: faTimes,
                    file: faFile,
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
                let formData = new FormData();
                this.files.forEach(f => {
                    formData.append('files', f, f.name)
                });

                axios.post("/api/analyze", formData, {headers: {'Content-Type': 'multipart/form-data'}})
                    .then(r => console.log(r.data))
                    .catch(e => console.log(e));

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