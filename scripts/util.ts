import util = require('util');
const generateKeyPair = util.promisify(require('crypto').generateKeyPair);

export const generateSignKeys = async () => {
    console.log("Generate Sign Keys:");

    try {
        // Generate random sign keys
        const { publicKey, privateKey } = await generateKeyPair('ed25519', {
            publicKeyEncoding: {
                type: 'spki',
                format: 'der'
            },
            privateKeyEncoding: {
                type: 'pkcs8',
                format: 'der',
            }
        });

        const keyPair = {
            publicKey: publicKey.toString('hex').substring(24),
            secretKey: privateKey.toString('hex').substring(32)
        };

        console.log(keyPair);

        return keyPair;

    } catch (err) {
        console.error("generateSignKeys error: ", err);
        throw err; // Re-throw error for handling at the caller's end
    }
}
