<template>
    <div v-show="waiting" class="waiting-block">
        <font-awesome-icon :icon="icons.waiting" size="2x" spin/>
        <span>Please wait...</span>
    </div>
</template>
<script>
    import {faSpinner} from '@fortawesome/free-solid-svg-icons'
    import {mapGetters} from 'vuex';

    export default {
        data() {
            return {
                icons: {
                    waiting: faSpinner,
                }
            }
        },
        computed: {
            ...mapGetters(['waiting'])
        },
    }
</script>
<style lang="scss">
    .waiting-block {
        color: #848bad;
        display: block;
        text-align: center;
        margin: 70px;
        @extend %mid-align;
        span {
            padding: 10px;
            display: inline-block;
            vertical-align: top;
            line-height: 0.7em;
            font-size: 1.1em;
            font-weight: bold;
        }
    }
</style>