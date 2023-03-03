# badger
Spammer notification service

# Requirements

      * Install Go 1.16+ 
      * Update slack url and channel value on config.json file 
      <img width="683" alt="image" src="https://user-images.githubusercontent.com/74706264/222638733-19c5f472-2ec1-40ff-b751-ea476b583911.png">
      
# Running Locally
      
    # step 1 : please run below cmd inside project 
       
       cmd 1 : go mod vendor
     
       cmd 2 : go run main.go 

       if any dependency error comes please install respective package using 
        
       go get <package name> 
       
       then 

       go mod vendor 
         
            or 

       go mod tidy

    # step 2 : server will start listen on like below 

         http://localhost:8000/


