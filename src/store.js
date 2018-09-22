import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        files: [],
        analyzedData: {}
    },
    getters: {
        files(s) {
            return s.files;
        },
        analyzedData(s) {
            return s.analyzedData;
        }
    },
    mutations: {
        addFiles(s, p) {
            let fileList = [];
            if (p.files || p.files.length) {
                fileList.push(...p.files)
            }
            p.value = "";
            fileList.forEach(fl => {
                if (!s.files.find(f => f.name === fl.name)) {
                    s.files.push(fl);
                }
            });
        },
        setAnalyzedData(s, p) {
            s.analyzedData = p;
        },
        removeFiles(s) {
            s.files = [];
        },
        removeFile(s, p) {
            s.files.splice(s.files.indexOf(p), 1);
        }
    },
    actions: {
        uploadFiles({commit, getters}) {
            let formData = new FormData();
            getters.files.forEach(f => {
                formData.append('files', f, f.name)
            });

            axios.post("/api/analyze", formData, {headers: {'Content-Type': 'multipart/form-data'}})
                .then(r => {
                    console.log(r.data)
                    commit('setAnalyzedData', r.data);
                })
                .catch(e => console.log(e));
        }
    }
})
