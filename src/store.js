import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        indexFile: null,
        selectedFiles: [],
        modules: [],
        schemeColors: ["#2aeaf5", "#ffa836", "#a59bda", "#ff6eb2", "#00cc58", "#feff48", "#0cb79d", "#ffd9ed", "#f1f1f1", "#bc80bd", "#ccebc5", "#ffed6f"],
        scrutinyData: {
            data: null,
            inProgress: false,
            elapsedTime: 0,
        },
        graphs: [
            {title: "Graph No.1", component: null},
            {title: "Graph No.2", component: null},
            {title: "Graph No.3", component: null},
            {title: "Graph No.4", component: null},
            {title: "Graph No.5", component: null}],
    },
    getters: {
        indexFile: s => s.indexFile,
        selectedFiles: s => s.selectedFiles,
        modules: s => s.modules,
        scrutinyData: s => s.scrutinyData,
        graphs: s => s.graphs,
    },
    mutations: {
        addSelectedFiles(s, fs) {
            let fileList = [];
            if (fs.files || fs.files.length) {
                fileList.push(...fs.files)
            }
            fs.value = "";
            let module = createModule();
            fileList.forEach(fi => {
                if (!s.selectedFiles.find(f => f.name === fi.name)) {
                    if (fi.type === "text/html") {
                        fi.module = module;
                    }
                    s.selectedFiles.push(fi);
                }
            });
            if (s.selectedFiles.some(sf => sf.module && (sf.module.name === module.name))) {
                s.modules.push(module);
            }

            function createModule() {
                let moduleName = "module-" + (s.modules.length + 1);
                if (s.modules.find(m => m.name === moduleName)) {
                    moduleName = moduleName + "-1"
                }
                let module = {
                    name: moduleName,
                    color: s.schemeColors.find(sc => !s.modules.map(m => m.color).includes(sc)),
                    editMode: false
                };
                return module
            }
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
        removeFile(s, file) {
            s.selectedFiles.splice(s.selectedFiles.indexOf(file), 1);
            if (!s.selectedFiles.some(sf => sf.module && (sf.module.name === file.module.name))) {
                s.modules.splice(s.modules.indexOf(file.module), 1)
            }
        },
        removeAllFiles(s) {
            s.indexFile = null;
            s.modules = [];
            s.selectedFiles = [];

            s.scrutinyData.elapsedTime = 0;
            s.scrutinyData.inProgress = false;
            s.scrutinyData.data = null;
        },
        removeModule(s, m) {
            let mf = s.selectedFiles
                .filter(sf => sf.module && (sf.module.name === m.name));
            mf.forEach(f => {
                let inx = s.selectedFiles.indexOf(f);
                s.selectedFiles.splice(inx, 1)
            });

            s.modules.splice(s.modules.indexOf(m), 1)
        },

        setScrutinyInProgress(s) {
            s.scrutinyData.inProgress = true;
        },
        setScrutinyElapsedTime(s, time) {
            s.scrutinyData.elapsedTime = time;
        },
        setScrutinyData(s, data) {
            s.scrutinyData.inProgress = false;
            s.scrutinyData.data = data;
        },
    },
    actions: {
        scrutinize({commit, getters}) {
            let startTime = new Date();
            commit('setScrutinyInProgress');
            commit('setScrutinyElapsedTime', 0);

            let formData = new FormData();
            let modules = [];

            getters.selectedFiles
                .map(sf => {
                    let file = Vue.util.extend({}, sf);
                    if (file.module) {
                        file.module = file.module.name;
                    }
                    return file;
                })
                .forEach(f => {
                    formData.append("files", f, f.name)
                });
            setTimeout(function () {
                let endTime = new Date();

                let elapsedTime = Math.round((endTime.getTime() - startTime.getTime()) / 100) / 10;

                commit('setScrutinyElapsedTime', elapsedTime);
                commit('setScrutinyData', {});
            }, 2350)
            // axios.post("/api/analyze", formData, {headers: {'Content-Type': 'multipart/form-data'}})
            //     .then(r => {
            //         //commit('setGraphData', r.data);
            //     })
            //     .catch(e => console.log(e))
            //     .finally(() => {
            //     });
        }
    }
})
