<template>
  <div class="home">
      <ErrorPage/>
    </div>
</template>

<script>
// @ is an alias to /src
import ErrorPage from '@/components/ErrorPage.vue'

export default {
  name: 'ErrorPage',
  components: {
    ErrorPage
  }
}
</script>
