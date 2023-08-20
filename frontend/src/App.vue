<script lang="ts" setup>
import {MediaDown} from '../wailsjs/go/cmd/Cmd'
import {ref} from "vue";
import { Message } from '@arco-design/web-vue';

const loading = ref(false)
const mediaUrl = ref("")
const msg = ref("")

async function down() {
  loading.value = true
  await MediaDown(mediaUrl.value)
  loading.value = false
  mediaUrl.value = ""
  Message.success("下载完成")
}
</script>
<template>
  <a-spin :loading="loading" tip="下载中...">
    <a-card>
      <a-form-item label="资源地址">
        <a-space>
          <a-input :style="{width:'320px'}" placeholder="请输入媒体地址进行下载！" v-model="mediaUrl" allow-clear/>
          <a-button type="primary" @click="down">下载</a-button>
        </a-space>
      </a-form-item>
      <!--<a-alert type="success" :show-icon="false">下载完成</a-alert>-->
    </a-card>
  </a-spin>
</template>

<style>
</style>
