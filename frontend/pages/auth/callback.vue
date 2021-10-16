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
    const requestBody = JSON.stringify(request.body())

    try {
      const response = await globalThis.fetch("http://localhost:8000/auth/callback", {
        method: "POST",
        mode: "cors",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      })
      const responseBody: CallbackResponse = await response.json()
      const token = responseBody.token
      console.log(token)
    }
    catch (e) {
      console.log(e)
    }
  },
  fetchOnServer: false,
}
</script>
