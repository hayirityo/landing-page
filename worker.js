// push 03/09/2026 00:09:10
export default {
  async fetch(){
    return new Response("landing-page", {
      headers: {"content-type":"text/plain"}
    })
  }
}
