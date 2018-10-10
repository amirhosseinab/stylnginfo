<template>
    <div class="user-guide-pane">
        <div v-for="(gl,index) in guideLines" class="user-guide-item" :key="index">
            <div v-if="!gl.checked" class="circle">
                {{index + 1 }}
            </div>
            <div v-else class="circle">
                <font-awesome-icon :icon="icons.done" size="lg" class="done"/>
            </div>
            <div class="guide-line-text">
                {{gl.guide}}
            </div>
        </div>
    </div>
</template>

<script>
    import {faCheck} from "@fortawesome/free-solid-svg-icons";

    export default {
        name: "v-user-guide",
        data() {
            return {
                guideLines: [
                    {guide: "Select index.html file from Files toolbox section", checked: false },
                    {guide: "Select HTML files and rename its module name", checked: false},
                    {guide: "Select CSS files", checked: false},
                    {guide: "Analyze!", checked: false},
                ],
                icons: {
                    done: faCheck,
                }
            }
        },
    }
</script>

<style scoped lang="scss">
    .user-guide-pane {
        align-self: center;
        opacity: .6;
        display: flex;
        flex-flow: column wrap;
        align-items: flex-start;
        justify-content: center;
        .user-guide-item {
            display: flex;
            flex-flow: row wrap;
            align-items: center;
            justify-content: flex-start;
            margin: .5rem;
            .circle {
                display: flex;
                align-items: center;
                justify-content: center;
                height: 3rem;
                width: 3rem;
                background-color: #cae4cb;
                color: #70b35d;
                border-radius: 100%;
                margin: 0 1rem;
                padding: .2rem;
                font-weight: bold;
                font-size: 1.2rem;
                .done {
                    padding: .3rem;
                }
            }
            .guide-line-text {
                flex: 1 0;
                color: $mid-gray-color;
                font-weight: bold;
                font-size: 1.4rem;
            }
        }

    }
</style>