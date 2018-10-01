import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        waiting: false,
        // files: [{name: 'cfp-list.html', type: 'text/html'}],
        files: [],
        // graphData: {
        //     htmlFiles: [{name: 'cfp-list.html', fileType: 0, selected: true}, {
        //         name: 'new-route-modal.html',
        //         fileType: 0,
        //         selected: true
        //     }, {name: 'select-flight-modal.html', fileType: 0, selected: true}, {
        //         name: 'add-or-edit-route-modal.html',
        //         fileType: 0,
        //         selected: true
        //     }],
        //     cssFiles: [{name: 'app.css', fileType: 1, selected: true}, {
        //         name: 'chat.css',
        //         fileType: 1,
        //         selected: true
        //     }, {name: 'ai.css', fileType: 1, selected: true}, {
        //         name: 'cfp.css',
        //         fileType: 1,
        //         selected: true
        //     }, {name: 'flight-search.css', fileType: 1, selected: true}]
        // },
        graphData: null,
    },
    getters: {
        waiting: s => s.waiting,
        analyzed: s => !!s.graphData,
        files: s => s.files,
        graphData: s => s.graphData,
        htmlFiles: s => s.graphData ? s.graphData.htmlFiles : [],
        cssFiles: s => s.graphData ? s.graphData.cssFiles : [],
        fileLinks: s => s.graphData ? s.graphData.fileLinks : [],
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
            let allFiles = s.graphData.htmlFiles.concat(s.graphData.cssFiles);
            if (allFiles.every(f => f.selected)) {
                allFiles.forEach(f => f.selected = false)
            } else {
                allFiles.forEach(f => f.selected = true)
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
