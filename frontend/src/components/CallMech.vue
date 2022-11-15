
<template>
  <div class="pa-5">
    <div id="map"></div>
    <v-card
      variant="tonal"
      @click="locatorButtonPressed()"
      class="text-left pa-4 d-flex justify-left align-center"
    >
      <v-icon icon="mdi-pin"></v-icon>
      <section>
        <v-card-title>Current Location</v-card-title>
        <v-card-text>{{currentLocation.lat+', '+currentLocation.lng}}</v-card-text>
      </section>
    </v-card>
    <v-spacer class="my-5"></v-spacer>

    <v-card variant="tonal" class="text-left pa-4" @click="selectCarModal=true">
      <v-card-title>Car detail</v-card-title>
      <section>
        <v-card-text>
          Type {{car.type}}, Brand {{car.brand}}
          <br />
          License plate: {{car.plate}}
        </v-card-text>
      </section>
    </v-card>

    <v-dialog v-model="selectCarModal">
      <v-card>
        <v-card-text v-for="(car,index) in cars" :key="index" @click="selectCar(car)">
          {{car.brand}} : {{car.plate}}
        </v-card-text>
        <v-card-text @click="selectCar(car)">
          {{car.brand}} : {{car.plate}}
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-spacer class="my-5"></v-spacer>
    <v-card variant="tonal" class="text-left pa-4">
      <v-card-title>Problem {{problem}}</v-card-title>
      <v-card-text class="d-flex">
        <v-list>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox v-model="problem" value="1" label="Problem1"></v-checkbox>
          </v-list-item>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox v-model="problem" value="3" label="Problem3"></v-checkbox>
          </v-list-item>
        </v-list>

        <v-list>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox v-model="problem" value="2" label="Problem2"></v-checkbox>
          </v-list-item>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox v-model="problem" value="4" label="Problem4"></v-checkbox>
          </v-list-item>
        </v-list>
      </v-card-text>
    </v-card>
    <v-spacer class="my-5"></v-spacer>
    <v-container class="pa-0 ma-0" fluid>
      <v-textarea
        outlined
        name="description"
        v-model="description"
        rows="5"
        label="Description"
        bg-color="white"
      ></v-textarea>
    </v-container>
    <v-spacer class="my-5"></v-spacer>
    <v-file-input counter v-model="files" multiple show-size truncate-length="20"></v-file-input>
    <v-spacer class="my-5"></v-spacer>
    <section class="text-center">
      <router-link to="/loading" class="text-decoration-none"><v-btn variant="tonal" color="green">Find Service</v-btn></router-link>
    </section>
  </div>
</template>

<script>
/* eslint-disable */
export default {
  data() {
    return {
      map: null,
      currentLocation: { lat: 0, lng: 0 },
      car: { id:"1", type: "SUV", brand: "MG", plate: "ก2113" },
      cars: [
        { id:"1", type: "SUV", brand: "MG", plate: "ก2113" },
        { id:"2", type: "Sedan", brand: "MG", plate: "ก4113" },
        { id:"3", type: "Van", brand: "MG", plate: "ก8113" }
      ],
      selectCarModal: false,
      problem: [],
      description: undefined,
      files: []
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
    },
    selectCar(car){
      this.car = car
      this.selectCarModal = false
    }
  }
};
</script>