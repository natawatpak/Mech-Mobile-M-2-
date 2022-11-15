package main

// Handler is a wrapper around lambda and mux so that we can continue to use mux via lambda
// func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	log.Print(request)
// 	if len(request.Body) < 1 {
// 		return events.APIGatewayProxyResponse{}, QueryNameNotProvided
// 	}

// 	var params struct {
// 		Query         string                 `json:"query"`
// 		OperationName string                 `json:"operationName"`
// 		Variables     map[string]interface{} `json:"variables"`
// 	}
// 	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
// 		log.Print("Could not decode body", err)
// 	}

// 	response := mainSchema.Exec(ctx, params.Query, params.OperationName, params.Variables)
// 	responseJSON, err := json.Marshal(response)
// 	if err != nil {
// 		log.Print("Could not decode body")
// 	}

// 	return events.APIGatewayProxyResponse{
// 		Body:       string(responseJSON),
// 		StatusCode: 200,
// 	}, nil
// }

// func main() {
// 	r := mux.NewRouter()

// 	var host string = "localhost"
// 	var port string = "5432"
// 	var user string = "postgres"
// 	var password string = "Eauu0244"
// 	var dbname string = "postgres"
// 	var goChiPort string = "8081"

// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	operator, err := resource.NewDBOperator(psqlInfo)
// 	util.CheckErr(err)
// 	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Not found", r.RequestURI)
// 		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
// 	})
// 	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: operator}})
// 	server := handler.NewDefaultServer(schema)
// 	r.Handle("/query", server)
// 	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	if runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtime_api != "" {
// 		adapter := gorillamux.NewV2(r)
// 		lambda.Start(adapter.ProxyWithContext)
// 	} else {
// 		srv := handler.NewDefaultServer(
// 			generated.NewExecutableSchema(
// 				generated.Config{
// 					Resolvers: &graph.Resolver{
// 						DB: operator,
// 					},
// 				},
// 			),
// 		)

// 		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 		r.Handle("/query", srv)

// 		log.Printf("connect to http://%s:%s/ for GraphQL playground", host, goChiPort)
// 		log.Fatal(http.ListenAndServe(":"+goChiPort, r))
// 	}
// }
