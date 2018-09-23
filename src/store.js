import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        waiting: false,
        files: [],
        analyzedData: {}
    },
    getters: {
        waiting: s => s.waiting,
        files: s => s.files,
        analyzedData: s => s.analyzedData,
        relatedCssFiles: s => htmlFileName => {
            var hf = s.analyzedData.htmlFiles.find(h => h.name === htmlFileName);
            if (!hf) return [];
            return [...new Set(hf.appliedSelectors.map(sel => sel.cssFileName))]
        },
        relatedSelectors: s => (htmlFileName, cssFileName) => {
            var hf = s.analyzedData.htmlFiles.find(h => h.name === htmlFileName);
            if (!hf) return [];

            var selectors = hf.appliedSelectors.filter(s => s.cssFileName === cssFileName);
            if (!selectors) return [];

            var fs = selectors.reduce((acc, c) => {
                acc.push(c.selectorName);
                return acc;
            }, []);
            return [...new Set(fs)]
        },
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
                    s.files.push(fl);
                }
            });
        },
        setAnalyzedData(s, p) {
            s.analyzedData = p;
        },
        removeFiles(s) {
            if (s.waiting) return;

            s.files = [];
            s.analyzedData = {};
        },
        removeFile(s, p) {
            if (s.waiting) return;
            s.files.splice(s.files.indexOf(p), 1);
        },
        wait(s, p) {
            s.waiting = p;
        },
    },
    actions: {
        uploadFiles({commit, getters}) {
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
