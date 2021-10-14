<template>
  <p></p>
</template>

<script lang="ts">
import Vue from 'vue'

class CallbackRequest {
  private state: string | null
  private code: string | null

  constructor(state: string | (string | null)[], code: string | (string | null)[]) {
    this.state = null
    this.code = null

    if (typeof state === "string") {
      this.state = state
    }
    if (typeof code === "string") {
      this.code = code
    }
  }

  isValid(): boolean {
    return (typeof this.state === "string") && (typeof this.code === "string")
  }

  body(): object {
    return {
      state: this.state,
      code: this.code,
    }
  }
}

interface CallbackResponse {
  token: string,
}

export default {
  async fetch() {
    const query = this.$nuxt.$route.query
    const request = new CallbackRequest(query.state, query.code)
    if (!request.isValid()) {
      console.log("invalid parameters")
      return
    }

    try {
      const response: CallbackResponse = await this.$http.$post("http://localhost:8000/auth/callback", request.body())
      const token = response.token
      console.log(token)
    }
    catch (e) {
      console.log(e)
    }
  },
  fetchOnServer: false,
}
</script>
