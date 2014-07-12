# A WebSocket chat with Go

Goal: make something to get started with Go.


Heroku Deployement
---------------------
  + Make sure you login to heroku 
    
    `heroku login`

  + Enable websockets
    
    `heroku labs:enable websockets`

  + Initialize git repo
    
    `git init`

    `git add -A .`

    `git commit -m "initial commit"`

  + Create Procfile
    
    `echo 'webapp: websocket_chat' > Profile`


  + Install Godep

    `go get github.com/kr/godep`

  + Save your dependencies

    `godep save`

  + Commit
    
    `git add -A .`

    `git commit -m "dependencies"`

  + Deploy to Heroku using buildpack

    `heroku create -b https://github.com/kr/heroku-buildpack-go.git`

  + Push to Heroku

    `git push heroku master`

  + If for some reason you need to stop the app

    `heroku ps:scale web=0`




