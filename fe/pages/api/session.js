import { withSessionRoute } from "../../lib/config/withSession";
export default withSessionRoute(createSessionRoute);

async function createSessionRoute(req, res) {
    if (req.method === "POST") {
        const { email, password } = req.body;
        // this will be the part where I contact the server
        if (email === VALID_EMAIL && password === VALID_PASSWORD) {
            req.session.user = {
                username: "test@gmail.com",
                isAdmin: true
            };
            await req.session.save();
            res.send({ ok: true });
        }
        return res.status(403).send("");
    }
    return res.status(404).send("");
}