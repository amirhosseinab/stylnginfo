import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        waiting: false,
        files: [],
        analyzedData: null,
    },
    getters: {
        waiting: s => s.waiting,
        analyzed: s => !!s.analyzedData,
        files: s => s.files,
        filteredFiles: s => s.filteredFiles,
        analyzedData: s => s.analyzedData,
        htmlFileNames: s => {
            return s.files.filter(f => f.type === 'text/html').map(f => f.name) || []
        },
        cssFileNames: s => {
            return s.files.filter(f => f.type === 'text/css').map(f => f.name) || []
        },
        htmlFiles: s => {
            if (!s.analyzedData) return [];
            return s.analyzedData.htmlFiles;
        }
    },
    mutations: {
        addFiles(s, p) {
            if (s.waiting) return;
            let fileList = [];
            if (p.files || p.files.length) {
                fileList.push(...p.files)
            }
            p.value = "";
            fileList.forEach(fl => {
                if (!s.files.find(f => f.name === fl.name)) {
                    s.files.push({file: fl, selected: true});
                }
            });
        },
        setAnalyzedData(s, p) {
            s.analyzedData = p;
        },
        removeFiles(s) {
            if (s.waiting) return;

            s.files = [];
            s.analyzedData = null;
        },
        removeFile(s, p) {
            if (s.waiting) return;
            s.files.splice(s.files.indexOf(p), 1);
        },
        // selectSelectors(s, p) {
        //     s.analyzedData.htmlFiles
        //         .find(hf => hf.name === p.htmlFileName).cssFiles
        //         .forEach(cf => {
        //             cf.selectors
        //                 .filter(s => s.name === p.selector)
        //                 .forEach(s => {
        //                     s.selected = !s.selected
        //                 })
        //         });
        // },
        toggleFileSelection(s, p) {
            p.selected = !p.selected;
        },
        toggleAllFiles(s) {

            if (s.files.every(f => f.selected)) {
                s.files.forEach(f => f.selected = false)
            } else {
                s.files.forEach(f => f.selected = true)
            }
        },
        wait(s, p) {
            s.waiting = p;
        },
    },
    actions: {
        analyzeFiles({commit, getters}) {
            if (getters.waiting) return;
            commit('wait', true);

            let formData = new FormData();
            getters.files.forEach(f => {
                formData.append('files', f, f.name)
            });

            axios.post("/api/analyze", formData, {headers: {'Content-Type': 'multipart/form-data'}})
                .then(r => {
                    commit('setAnalyzedData', r.data);
                })
                .catch(e => console.log(e))
                .finally(() => commit('wait', false));
        }
    }
})
