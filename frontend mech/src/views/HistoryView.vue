<template>
  <div class="order">
    <HistoryPage/>
  </div>
</template>

<script>
// @ is an alias to /src
import HistoryPage from '@/components/HistoryPage.vue'

export default {
  name: 'HistoryView',
  components: {
    HistoryPage
  }
}
</script>
