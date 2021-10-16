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
    const response = await globalThis.fetch("http://localhost:8000/auth", {
      method: "POST",
      mode: "cors",
      credentials: "include",
    })
    const body: AuthStartResponse = await response.json()
    const authorizationUri = body.authorizationUri
    window.location.assign(authorizationUri)
  },
  fetchOnServer: false,
}
</script>
