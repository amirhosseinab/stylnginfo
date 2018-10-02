import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        waiting: false,
        files: [],
        graphData: null,
    },
    getters: {
        waiting: s => s.waiting,
        analyzed: s => !!s.graphData,
        files: s => s.files,
        graphData: s => s.graphData,
        htmlFiles: s => s.graphData ? s.graphData.htmlFiles : [],
        cssFiles: s => s.graphData ? s.graphData.cssFiles : [],
    },
    mutations: {
        handleFiles(s, p) {
            if (s.waiting) return;
            let fileList = [];
            if (p.files || p.files.length) {
                fileList.push(...p.files)
            }
            p.value = "";
            fileList.forEach(fi => {
                if (!s.files.find(f => f.name === fi.name)) {
                    s.files.push(fi);
                }
            });
        },
        setGraphData(s, p) {
            p.htmlFiles ? (p.htmlFiles.forEach(hf => hf.selected = true)) : p.htmlFiles;
            p.cssFiles ? (p.cssFiles.forEach(cf => cf.selected = true)) : p.cssFiles;
            s.graphData = p;
        },
        renew(s) {
            if (s.waiting) return;

            s.files = [];
            s.graphData = null;
        },
        removeFile(s, p) {
            if (s.waiting) return;
            s.files.splice(s.files.indexOf(p), 1);
        },
        toggleFileSelection(s, p) {
            p.selected = !p.selected;
        },
        toggleAllFilesSelection(s) {
            let files = s.graphData.htmlFiles.concat(s.graphData.cssFiles);
            if (files.every(f => f.selected)) {
                files.forEach(f => f.selected = false)
            } else {
                files.forEach(f => f.selected = true)
            }
        },
        toggleHtmlFilesSelection(s) {
            let files = s.graphData.htmlFiles;
            if (files.every(f => f.selected)) {
                files.forEach(f => f.selected = false)
            } else {
                files.forEach(f => f.selected = true)
            }
        },
        toggleCssFilesSelection(s) {
            let files = s.graphData.cssFiles;
            if (files.every(f => f.selected)) {
                files.forEach(f => f.selected = false)
            } else {
                files.forEach(f => f.selected = true)
            }
        },
        wait(s, p) {
            s.waiting = p;
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
