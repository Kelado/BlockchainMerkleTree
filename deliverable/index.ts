import { keccak_256 } from "@noble/hashes/sha3";
import { MerkleTree } from "merkletreejs";

// Solution 1
// You need to use ALLOWLIST length equals to the power of 2
// e.g. 2,4,8,16 ...

// Solution 2
// You can have as many as you want
const ALLOWLIST = [
	"sei1vwap3mwdddyrcenn4f3fytd3rvxp7tp08eskp8",
	"sei1vtpy9gu4xtaddzxqxalj3z02zt53hzdcstpy77",
	"sei15ggkgs9cth485n299cftf3v939m85ludgms9pu",
	"sei1g9zyz302ln5nzxyca6fqsre9qnrp7khwywtqf7",
	"sei1msqmz3euz4hwdvjn22fckfry6tvwk264u2fg0t",
	"sei10wx2t6cfcz8qaq77a2dx7fc74pw3ds3f5ssv20",
	"sei1s772ntuxfz2w02jwu8z9lnqcmqzfs7uhgpeda7",
	"sei13kawxavfda2w8f8sh2xcr92u6yxjdgjce0qvgs",
	"sei1rdd6v2ufrf5h0qm3edc69l6g6xsz6l07z3wtd5",
	"sei1x2r5r5ap7gnyxqwc3dpp8603aqyh23athy2m9d",
];

const WALLET_ADDRESS = "sei1vwap3mwdddyrcenn4f3fytd3rvxp7tp08eskp8";

(async () => {

	const hashedWallets = ALLOWLIST.map(keccak_256);

	let options = {
		sortPairs: true,
		duplicateOdd: true // Solution 2, follow same procedure as golang script
	}

	const tree = new MerkleTree(hashedWallets, keccak_256, options);

	const merkleRoot = tree.getRoot().toString("hex");
	const merkleProof = tree
		.getProof(Buffer.from(keccak_256(WALLET_ADDRESS)))
		.map((element) => Array.from(element.data));

	console.log(merkleRoot);
	console.log(merkleProof);
})();
