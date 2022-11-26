<template>
  <div class="text-left ma-4">
    <div>
      <v-alert
        color="teal"
        title="Incoming order"
        variant="tonal"
        icon="mdi-ballot"
      >
        <div class="d-flex flex-row align-center justify-space-between">
          <div>
            There was an incoming order from your client. Go check now!
          </div>

          <v-btn
            color="teal-darken-1"
            variant="outlined"
          >
            Go
          </v-btn>
        </div>
      </v-alert>
    </div>

    <div class="mt-4">
      <v-text class="text-h5">Today summary</v-text>
    </div>

    <v-row no-gutters>
      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>Total order</v-card-title>
          <v-card-text class="text-center text-h6">{{total}}</v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>On the way</v-card-title>
          <v-card-text class="text-center text-h6">{{otw}}</v-card-text>
        </v-card>
      </v-col>

      <v-responsive width="100%"></v-responsive>

      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>In process</v-card-title>
          <v-card-text class="text-center text-h6">{{onprocess}}</v-card-text>
        </v-card>
      </v-col>

      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>Finish</v-card-title>
          <v-card-text class="text-center text-h6">{{finish}}</v-card-text>
        </v-card>
      </v-col>
    </v-row>


        <v-tabs v-model="tab" grow class="ma-4">
          <v-tab class="text-h5" value="one">On process</v-tab>
          <v-tab class="text-h5" value="two">Finish</v-tab>
        </v-tabs>

        <v-card-text>
          <v-window v-model="tab">
            <v-window-item value="one">
              <v-content>
                <v-container v-for="item in items" :key="item.id" class="pa-0 my-4">
                  <v-card width="100%" variant="tonal" class="text-left pa-4 ">
                    <v-row class="pa-2">
                      <v-card-title>
                        Order {{ item.id }} | {{ item.username }}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{
                        item.status
                      }}</v-chip>
                    </v-row>
                    <v-card-text class="text-h7">
                      Car {{ items.type }} | {{ item.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>

            <v-window-item value="two">
              <v-card variant="tonal" class="text-left my-4 pa-4">
                <v-row class="pa-2">
                  <v-card-title
                    >Order {{ details.id }} | {{ details.name }}</v-card-title
                  >
                  <v-chip class="ma-2" color="green">{{details.status}}</v-chip>
                </v-row>
                <section>
                  <v-card-text class="text-h7">
                    Car {{ details.type }} | {{ details.brand }} <br />
                    Problems {{ details.problem }} <br />
                    {{ details.location }} km
                  </v-card-text>
                </section>
              </v-card>
            </v-window-item>
          </v-window>
        </v-card-text>
  </div>
</template>

<style>
.text-h4 {
  display: flex;
  justify-content: left;
  align-content: center;
}
</style>

<script>
export default {
  data() {
    return {
      dialog: false,
      orders: 1,
      total: 5,
      otw: 1,
      onprocess: 1,
      finish: 1,
      problems: ["no battery", "broken motor"],
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.",
      location: {
        lat: "20",
        lng: "90"
      },
      details: {
        id: "12345",
        type: "SUV",
        brand: "MG",
        plate: "ศง 4727",
        name: "nunnapat",
        problem: "2",
        location: "2.1",
        status: "Finish",
      },
      items: [
        {
          id: "11111",
          username: "nunnapat",
          distance: "2.2",
          type: "SUV",
          brand: "MG",
          problem: "2",
          status: "On-process",
        },
        {
          id: "22222",
          username: "natawat",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "On the way",
        },
        {
          id: "33333",
          username: "kiitiphum",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          problem: "2",
          status: "Accept",
        },
      ],
      tab: null,
    };
  },
};
</script>