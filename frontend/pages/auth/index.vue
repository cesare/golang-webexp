<template>
  <p>test for echo</p>
</template>

<script lang="ts">
import { NuxtHTTPInstance } from '@nuxt/http'
import Vue from 'vue'

interface AuthStartResponse {
  authorizationUri: string,
}

export default {
  async fetch() {
    const response: AuthStartResponse = await this.$http.$post("http://localhost:8000/auth")
    const authorizationUri = response.authorizationUri
    window.location.assign(authorizationUri)
  },
  fetchOnServer: false,
}
</script>
