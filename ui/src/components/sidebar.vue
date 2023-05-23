

<template>
    <div class="sider">
        <div class="header">
            <div class="avatar">
                <img src="../assets/avatar.png">
            </div>

        </div>
        <div class="tag">
            <div class="tag-item">
                <div class="title">文章</div>
                <div class="number">{{ article_total }}</div>
            </div>
            <div class="tag-item">
                <div class="title">栏目</div>
                <div class="number">{{ category_total }}</div>
            </div>
            <div class="tag-item">
                <div class="title">标签</div>
                <div class="number">{{ subCategory_total }}</div>
            </div>
        </div>
        <div class="menu">
            <div class="menu-item-parent" @click="toLink(0, 0)">
                <div class="menu-item" :class="active.join() == [0, 0].join() ? 'active' : ''">
                    <div class="menu-icon">
                        <img src="../assets/house.svg" alt="" srcset="" style="width: 24px;">
                    </div>
                    <div class="menu-content">首页</div>

                    <div style="width: 17px;"></div>
                </div>
            </div>
            <div v-for="item in menu" :key="item.id" @click="status = item.id">
                <div class="menu-item-parent">
                    <div class="menu-item">
                        <div class="menu-icon">
                            <img :src="'https://img.shields.io/badge/' + item.name + '-000000?logo=' + item.name + '&logoColor=FADC6A'" alt="" srcset="" style="width: 45px;">
                        </div>
                        <div class="menu-content">{{ item.name }}</div>
                        <div class="underline"
                            :class="status == item.id ? 'menu-item-parent-active' : 'menu-item-parent-disabled'">
                        </div>
                    </div>
                </div>
                <div class="menu-item-children" :class="status == item.id ? 'menu-item-active' : 'menu-item-disabled'"
                    v-for="child in item.children" :key="child.name" @click="toLink(item.id, child.id)">
                    <div class="menu-item" :class="active.join() == [item.id, child.id].join() ? 'active' : ''">
                        <div class="menu-icon">
                            <img :src="'https://img.shields.io/badge/' + child.name + '-000000?logo=' + child.name + '&logoColor=FADC6A'"
                                style="width:45px;">
                        </div>
                        <div class="menu-content">{{ child.name }}</div>
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>
<script setup>
import { ref, reactive } from 'vue';
let status = ref(0)
let active = reactive([0, 0])
const emit = defineEmits(['switch_article']);

const props = defineProps({
    article_total: {
        type: Number,
        default: 0,
    },
    category_total: {
        type: Number,
        default: 0,
    },
    subCategory_total: {
        type: Number,
        default: 0,
    },
    menu: { type: Array, default: () => [] }
});
function toLink(article_id, sub_category_id) {
    active[0] = article_id
    active[1] = sub_category_id
    emit('switch_article', article_id, sub_category_id);
}

</script>
<style lang="scss" scoped>
$sider-width: 320px;


.sider {
    height: 98vh;
    box-sizing: border-box;
    overflow: auto;
    width: $sider-width;
    padding: 30px;
    margin-top: 1vh;
    text-align: center;

    .avatar {
        border-radius: 100%;
        margin-bottom: 40px;

        img {
            border-radius: 100%;
            height: 160px;
            width: 160px;
            object-fit: cover;
            object-position: center;
        }
    }

    .tag {
        gap: 20px;
        display: flex;
        justify-content: center;
        margin-bottom: 20px;

        .tag-item {
            line-height: 30px;

            .title {
                font-size: 16px;
                font-weight: bold;
            }

            .number {
                font-size: 14px;
                color: #868e96;
            }
        }


    }

    .menu {
        .menu-item-parent {

            .menu-item {
                text-align: center;
                cursor: pointer;
                font-size: 18px;
                display: flex;
                align-items: center;
                justify-content: center;
                width: 200px;
                margin: 20px auto;
                padding: 5px;
                border-radius: 20px;
                border-radius: 10px;
                color: #868e96;
                background: rgba(0, 0, 0, .8);
                box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .5);

                &:hover {
                    color: white;
                    background: #fe5f77;
                }

                .menu-icon {
                    flex: 1;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                }

                .menu-content {
                    flex: 2;
                }

                .underline {
                    margin-right: 10px;
                    width: 10px;
                    height: 10px;
                    background-color: transparent;
                    /* 模块背景为透明 */
                    border-color: #f02c57;
                    border-style: solid;
                    border-width: 2px 2px 0 0;
                    transform: rotate(45deg);
                }
            }
        }

        .menu-item-parent-active {
            transform: rotate(135deg) !important;
            transition: all .5s ease-out;
        }

        .menu-item-parent-disabled {
            transition: all .5s ease-out;
            transform: rotate(45deg) !important;
        }

        .menu-item-children {
            position: absolute;
            opacity: 0;
            opacity: 0;
            transform: translateY(-100px);

            .menu-item {
                cursor: pointer;
                font-size: 18px;
                display: flex;
                align-items: center;
                justify-content: center;
                text-align: center;
                width: 180px;
                margin: 10px 0px 10px 40px;
                padding: 5px;
                border-radius: 20px;
                border-radius: 10px;
                color: #868e96;
                background: rgba(0, 0, 0, .8);
                box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .5);

                &:hover {
                    color: white;
                    background: #fe5f77;
                }

                .menu-icon {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    flex: 1;
                }

                .menu-content {
                    text-align: center;
                    white-space: nowrap;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    flex: 3;
                }

            }
        }

        .menu-item-active {

            position: relative;
            opacity: 1;
            visibility: visible;
            /*将元素设为可见*/
            transition: all 1s ease-out;
            transform: translateY(0px);
        }

        .menu-item-disabled {
            position: absolute;
            opacity: 0;
            visibility: hidden;
            transform: translateY(-100px);
            /*将元素设为可见*/
            transition: all .8s ease-out;
        }
    }
}

.active {
    color: white !important;
    background: #fe5f77 !important;
}
</style>
  