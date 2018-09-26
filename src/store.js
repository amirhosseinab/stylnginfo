import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        waiting: false,
        filter: {
            showCssFiles: true,
            showSelectors: true,
        },
        files: [],
        analyzedData: null,
    },
    getters: {
        waiting: s => s.waiting,
        analyzed: s => !!s.analyzedData,
        filter: s => s.filter,
        files: s => s.files,
        analyzedData: s => s.analyzedData,
        htmlFileNames: s => {
            return s.files.filter(f => f.type === 'text/html').map(f => f.name) || []
        },
        cssFileNames: s => {
            return s.files.filter(f => f.type === 'text/css').map(f => f.name) || []
        },
        htmlFiles: s => {
            if (!s.analyzedData) return [];
            return s.analyzedData.htmlFiles
                .filter(hf => s.files
                    .filter(f => f.selected)
                    .map(f => f.file.name).indexOf(hf.name) !== -1);
        },
        cssFiles: s => {
            if (!s.analyzedData) return [];
            return [...(new Set(s.analyzedData.htmlFiles.reduce((acc, hf) => {
                acc.push(...hf.cssFiles);
                return acc
            }, []).map(cf => cf.name)))];

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
                formData.append("files", f.file, f.file.name)
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
