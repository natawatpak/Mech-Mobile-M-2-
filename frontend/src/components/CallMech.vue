
<template>
  <div class="pa-5">
    <div id="map"></div>
    <v-card variant="tonal" class="text-left pa-4 d-flex justify-left align-center">
      <v-icon icon="mdi-pin"></v-icon>
      <section>
        <v-card-title>Current Location</v-card-title>
        <v-card-text>{{currentLocation.lat+', '+currentLocation.lng}}</v-card-text>
      </section>
    </v-card>
    <section class="text-center ma-4">
      <v-btn @click="locatorButtonPressed()">Use Current location</v-btn>
    </section>

    <v-card variant="tonal" class="text-left pa-4 ">
      <v-card-title>Car detail</v-card-title>
      <section>
        <v-card-text>Type {{car.type}}, Brand {{car.brand}} <br/> License plate: {{car.plate}}</v-card-text>
      </section>
    </v-card>

  </div>
</template>

<script>
/* eslint-disable */
export default {
  data() {
    return {
      map: null,
      currentLocation: { lat: 0, lng: 0 },
      car: {type:"SUV", brand:"MG", plate:"ก2113"}
    };
  },
  mounted() {
    this.initMap();
    this.locatorButtonPressed();
    this.setMarker(this.mapCenter, "A");
  },
  methods: {
    locatorButtonPressed() {
      const success = position => {
        const latitude = position.coords.latitude;
        const longitude = position.coords.longitude;
        // Do something with the position
        console.log(latitude + ", " + longitude);
        this.currentLocation.lat = latitude;
        this.currentLocation.lng = longitude;
        this.setMarker(this.currentLocation, "A");
      };

      const error = err => {
        console.log(error);
      };

      // This will open permission popup
      navigator.geolocation.getCurrentPosition(success, error);
    },
    initMap() {
      this.map = new google.maps.Map(document.getElementById("map"), {
        center: this.currentLocation,
        zoom: 5,
        maxZoom: 20,
        minZoom: 3,
        streetViewControl: false,
        mapTypeControl: false,
        fullscreenControl: true,
        zoomControl: true
      });
    },
    setMarker(Points, Label) {
      const markers = new google.maps.Marker({
        position: Points,
        map: this.map,
        label: {
          text: Label,
          color: "#FFF"
        }
      });
    }
  }
};
</script>