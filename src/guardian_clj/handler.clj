(ns guardian-clj.handler
  (:require [compojure.core :refer :all]
            [compojure.route :as route]
            [guardian-clj.database-handler :refer [create-db fixtures output]]
            [ring.middleware.defaults :refer [wrap-defaults site-defaults]]))

(defroutes app-routes
  (GET "/" [] "Hello World")
  (GET "/db" []
       (create-db)
       (fixtures)
       (let [result output]
         (keys (first result))
         (:body (first result))))
  (route/not-found "Not Found"))

(def app
  (wrap-defaults app-routes site-defaults))
