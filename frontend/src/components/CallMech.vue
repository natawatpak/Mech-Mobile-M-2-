<template>
  <div class="pa-5">
    <v-card text-left pa-4 d-flex justify-left align-center>
      <div style="width:auto;height:15rem;" id="map"></div>
    </v-card>
    <v-card
      variant="tonal"
      @click="locatorButtonPressed()"
      class="text-left pa-4 d-flex justify-left align-center"
    >
      <v-icon icon="mdi-pin"></v-icon>
      <section>
        <v-card-title>Current Location</v-card-title>
        <v-card-text>{{
          currentLocation.lat + ", " + currentLocation.lng
        }}</v-card-text>
      </section>
    </v-card>
    <v-spacer class="my-5"></v-spacer>

    <v-card variant="tonal" class="text-left pa-4">
      <v-row class="pl-5">
        <v-card-title>Car detail</v-card-title>
        <v-spacer></v-spacer>
        <v-btn class="mx-1" @click="selectCarModal = true"
          >choose from preset</v-btn
        >
      </v-row>
      <section>
        <v-card-text>
          Type {{ car.type }}, Brand {{ car.brand }}
          <br />
          License plate: {{ car.plate }}
        </v-card-text>
      </section>
    </v-card>

    <v-dialog v-model="selectCarModal">
      <v-card>
        <v-card-text
          v-for="(car, index) in cars"
          :key="index"
          @click="selectCar(car)"
        >
          {{ car.brand }} : {{ car.plate }}
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-spacer class="my-5"></v-spacer>
    <v-card variant="tonal" class="text-left pa-4">
      <v-card-title>Problems</v-card-title>
      <v-card-text class="d-flex">
        <v-list>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="1"
              label="Problem1"
            ></v-checkbox>
          </v-list-item>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="3"
              label="Problem3"
            ></v-checkbox>
          </v-list-item>
        </v-list>

        <v-list>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="2"
              label="Problem2"
            ></v-checkbox>
          </v-list-item>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="4"
              label="Problem4"
            ></v-checkbox>
          </v-list-item>
        </v-list>
      </v-card-text>
    </v-card>
    <v-spacer class="my-5"></v-spacer>
    <v-card variant="tonal" class="text-left pa-4">
      <v-card-title>Details</v-card-title>
      <v-textarea
        outlined
        name="description"
        v-model="description"
        rows="5"
        label="Description"
        bg-color="white"
      ></v-textarea>
    </v-card>
    <v-spacer class="my-5"></v-spacer>
    <v-card variant="tonal" class="text-left pa-4">
      <v-card-title>Upload files</v-card-title>
      <v-file-input
        counter
        v-model="files"
        accept="multiple show-size truncate-length=20"
        outlined
        label="Click here to attached photo or video"
      ></v-file-input>
    </v-card>
    <v-spacer class="my-5"></v-spacer>
    <section class="text-center">
      <router-link to="/loading" @click="addTicket"
        ><v-btn>Find Service</v-btn></router-link
      >
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
      car: { id: "1", type: "SUV", brand: "MG", plate: "ก2113" },
      cars: [],
      selectCarModal: false,
      problem: ["tyres", "flat"],
      description: undefined,
      files: [],
    };
  },
  mounted() {
    this.initMap();
    this.locatorButtonPressed();
    this.setMarker(this.mapCenter, "A");
    sessionStorage.setItem("cusID", "c57a987c-0c23-43fb-a131-c73f1e37d2fe");
    sessionStorage.setItem("fName", "phum");
    sessionStorage.setItem("lName", "kitiphum");
    sessionStorage.setItem("tel", "0123456789");
    sessionStorage.setItem("email", "phum@gmail.com");
    this.getCarList();
  },
  methods: {
    locatorButtonPressed() {
      const success = (position) => {
        const latitude = position.coords.latitude;
        const longitude = position.coords.longitude;
        // Do something with the position
        console.log(latitude + ", " + longitude);
        this.currentLocation.lat = latitude;
        this.currentLocation.lng = longitude;
        this.setMarker(this.currentLocation);
        this.map.setCenter(this.currentLocation);
      };

      const error = (err) => {
        console.log(error);
      };

      // This will open permission popup
      navigator.geolocation.getCurrentPosition(success, error);
    },
    initMap() {
      this.map = new google.maps.Map(document.getElementById("map"), {
        center: this.currentLocation,
        zoom: 15,
        maxZoom: 20,
        minZoom: 3,
        streetViewControl: false,
        mapTypeControl: false,
        fullscreenControl: false,
        zoomControl: true,
      });
    },
    setMarker(Points, Label) {
      const markers = new google.maps.Marker({
        position: Points,
        map: this.map,
        label: {
          text: Label,
          color: "#FFF",
        },
      });
    },
    selectCar(car) {
      this.car = car;
      this.selectCarModal = false;
    },
    getCarList() {
      const data = new URLSearchParams({
        cusID: sessionStorage.getItem("cusID"),
      });
      console.log(sessionStorage.getItem("cusID"));
      this.axios
        .post("http://localhost:3000/customer/get-car-list", data)
        .then((response) => {
          console.log(response.data);
          this.cars = response.data;
        });
    },
    addTicket() {
      const data = new URLSearchParams({
        cusID: sessionStorage.getItem("cusID"),
        carID: this.car.id,
        problem: this.problem,
        status: "Active",
      });
      this.axios
        .post("http://localhost:3000/customer/add-ticket", data)
        .then((response) => {
          console.log(response.data);
          sessionStorage.push("ticketID", response.data);
        });
    },
  },
};
</script>