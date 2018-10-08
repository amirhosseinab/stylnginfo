import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        selectedFiles: [],
        indexFile: null,
    },
    getters: {
        selectedFiles: s => s.selectedFiles,
        indexFile: s => s.indexFile,
    },
    mutations: {
        addSelectedFiles(s, fs) {
            let fileList = [];
            if (fs.files || fs.files.length) {
                fileList.push(...fs.files)
            }
            fs.value = "";
            fileList.forEach(fi => {
                if (!s.selectedFiles.find(f => f.name === fi.name)) {
                    s.selectedFiles.push(fi);
                }
            });
        },
        addIndexFile(s, fs) {
            if (fs.files || fs.files.length) {
                s.indexFile = fs.files[0]
            }
            fs.value = "";
        },
        removeIndexFile(s) {
            s.indexFile = null
        },
        removeFile(s,file) {
            s.selectedFiles.splice(s.selectedFiles.indexOf(file), 1)
        },
        removeAllFiles(s) {
            s.selectedFiles = [];
            s.indexFile = null;
        },
    },
    actions: {
        analyze({commit, getters}) {
            if (getters.waiting) return;
            commit('wait', true);

            let formData = new FormData();
            getters.files.forEach(f => {
                formData.append("files", f, f.name)
            });

            axios.post("/api/analyze", formData, {headers: {'Content-Type': 'multipart/form-data'}})
                .then(r => {
                    commit('setGraphData', r.data);
                })
                .catch(e => console.log(e))
                .finally(() => commit('wait', false));
        }
    }
})
