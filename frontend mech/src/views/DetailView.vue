<template>
  <div class="home">
    <DetailsPage/>
  </div>
</template>

<script>
// @ is an alias to /src
import DetailsPage from '@/components/OrderDetails.vue'

export default {
  name: 'DetailsView',
  components: {
    DetailsPage
  }
}
</script>
