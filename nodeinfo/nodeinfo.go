package nodeinfo

import "../cryptorsa"

var NodeID = "cusat"
var PrivateKey = cryptorsa.GetRsaPrivateKeyFromPem("/app/nodeinfo/privatekey/" + NodeID + ".pem")
var Difficulty = "0000"
