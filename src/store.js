import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        indexFile: null,
        selectedFiles: [],
        modules: [],
        schemeColors: ["#2aeaf5", "#ffa836", "#a59bda", "#ff6eb2", "#00cc58", "#feff48", "#0cb79d", "#ffd9ed", "#f1f1f1", "#bc80bd", "#ccebc5", "#ffed6f"],
        selectedGraphName: null,
        storedData: null,
        graphs: [
            {
                name: "css-files-weight",
                title: "CSS Files Weight",
                component: "v-graph-css-files-weight",
                data: {},
                inProgress: false,
                elapsedTime: null
            },
            // {
            //     name: "html-inline-style",
            //     title: "Inline Style Pollution",
            //     component: "v-graph-inline-style",
            //     data: {},
            //     inProgress: false,
            //     elapsedTime: null
            // },
        ],
    },
    getters: {
        indexFile: s => s.indexFile,
        selectedFiles: s => s.selectedFiles,
        modules: s => s.modules,
        selectedGraphName: s => s.selectedGraphName,
        selectedGraph: s => s.graphs.find(g => g.name === s.selectedGraphName),
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
        updateSelectedGraph(s, graphName) {
            if (s.graphs.find(g => g.name === graphName && Object.keys(g.data).length > 0)) {
                s.selectedGraphName = graphName;
            }
        },
        deselectAllGraphs(s) {
            s.selectedGraphName = null;
        },
        addIndexFile(s, fs) {
            if (fs.files || fs.files.length) {
                s.indexFile = fs.files[0];
            }
            fs.value = "";
        },
        removeIndexFile(s) {
            s.indexFile = null;
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
            s.outdatedGraphs = false;
            s.selectedGraphName = null;

            s.graphs.forEach(g => {
                g.inProgress = false;
                g.elapsedTime = null;
                g.data = {};
            })
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
        updateGraph(s, {graphName, data, elapsedTime}) {
            let graph = s.graphs.find(g => g.name === graphName);
            if (!graph) return;
            graph.data = data;
            graph.elapsedTime = elapsedTime;
        },
        setGraphState(s, {graphName, inProgress}) {
            let graph = s.graphs.find(g => g.name === graphName);
            if (!graph) return;
            graph.inProgress = inProgress;
        }
    },
    actions: {
        getCssFilesWeight({commit, getters}) {
            // This method provide data for a bubble chart of css files' weight
            // sample: https://beta.observablehq.com/@mbostock/d3-bubble-chart
            //  .map(sf => {
            //       let file = Vue.util.extend({}, sf);
            //       if (file.module) {
            //           file.module = file.module.name;
            //       }
            //       return file;
            //   })

            const graphName = "css-files-weight";
            const startTime = new Date();

            commit('setGraphState', {graphName, inProgress: true});

            let formData = new FormData();
            formData.append("indexFile", getters.indexFile, getters.indexFile.name);
            getters.selectedFiles
                .filter(sf => sf.type === "text/css")
                .forEach(f => {
                    formData.append("files", f, f.name)
                });

            axios.post("/api/CssFilesWeight", formData, {headers: {'Content-Type': 'multipart/form-data'}})
                .then(({data}) => {
                    let endTime = new Date();
                    let elapsedTime = Math.round((endTime.getTime() - startTime.getTime()) / 100) / 10;
                    commit('updateGraph', {graphName, elapsedTime, data});
                })
                .catch(e => console.log(e))
                .finally(() => commit('setGraphState', {graphName, inProgress: false}))
        },

        getHtmlCssRelations({commit, getters}) {
            // This method provide relations between HTML & CSS files and its
            // weight based on selector usage.
            // sample: https://bost.ocks.org/mike/uberdata/
        },
        scrutinize({commit, dispatch}) {
            commit("deselectAllGraphs");
            dispatch("getCssFilesWeight");
        }
    }
})
