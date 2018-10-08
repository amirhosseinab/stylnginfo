<template>
    <div :class="['btn',{'clicked':clicked}]" @mousedown="clicked=true" @mouseup="clicked=false"
         @click="$emit('click',$event)">
        <font-awesome-icon v-if="icon" class="btn-icon" :icon="icon"/>
        <div class="btn-text">{{title}}</div>
    </div>
</template>

<script>
    export default {
        name: "v-button",
        data() {
            return {
                clicked: false,
            }
        },
        props: {
            icon: {
                type: Object,
                required: false
            },
            title: {
                type: String,
                required: true
            }
        }
    }
</script>

<style scoped lang="scss">
    .btn {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: .6rem .6rem;
        margin: .2rem;
        background-color: transparent;
        border-radius: 8px;
        border: solid 1px $white-color;
        color: $white-color;
        cursor: pointer;
        transition: all .1s ease;
        transform: translate(-.03rem, -.05rem);
        box-shadow: 2px 3px 10px 1px lighten($dark-gray-color, 4%);
        font-size: .75rem;
        &.clicked {
            transform: translate(.03rem, .05rem);
            box-shadow: none;
        }
        .btn-text {
            padding: 0 .1rem;
            transform: translateY(.01rem);
        }
        .btn-icon {
            margin: 0 .3rem;
        }

    }
</style>