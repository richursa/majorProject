package nodeinfo

import "../cryptorsa"

var NodeID = "cusat"
var PrivateKey = cryptorsa.GetRsaPrivateKeyFromPem(NodeID + ".pem")
