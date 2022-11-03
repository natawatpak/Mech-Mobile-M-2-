
<template>
  <div class="text-center ma-4">
    <v-btn @click="locatorButtonPressed()">Get Current location</v-btn>
    <GMapMap :center="center" :zoom="7" map-type-id="terrain" style="width: 500px; height: 300px">
      <GMapCluster>
        <GMapMarker
          :key="index"
          v-for="(m, index) in markers"
          :position="m.position"
          :clickable="true"
          :draggable="true"
          @click="center=m.position"
        />
      </GMapCluster>
    </GMapMap>
  </div>
</template>

<script>
/* eslint-disable */
export default {
  data() {
    return {
      center: { lat: 51.093048, lng: 6.84212 },
      markers: [
        {
          position: {
            lat: 51.093048,
            lng: 6.84212
          }
        } // Along list of clusters
      ]
    };
  },
  methods: {
    locatorButtonPressed() {
      const success = position => {
        const latitude = position.coords.latitude;
        const longitude = position.coords.longitude;
        // Do something with the position
        console.log(latitude + ", " + longitude);
      };

      const error = err => {
        console.log(error);
      };

      // This will open permission popup
      navigator.geolocation.getCurrentPosition(success, error);
    }
  }
};
</script>